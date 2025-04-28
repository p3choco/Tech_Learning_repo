import os
import pytest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
from selenium.webdriver import Remote

@pytest.fixture(scope="session")
def driver():
    bs_user = os.getenv("BROWSERSTACK_USERNAME")
    bs_key  = os.getenv("BROWSERSTACK_ACCESS_KEY")

    options = Options()
    options.set_capability("browser", "Chrome")
    options.set_capability("browser_version", "latest")
    options.set_capability("os", "Windows")
    options.set_capability("os_version", "10")
    options.set_capability("name", "Selenium Python Test Suite")
    options.set_capability("build", "Selenium-Python-BrowserStack")
    options.set_capability("browserstack.local", "true")

    remote_url = f"https://{bs_user}:{bs_key}@hub-cloud.browserstack.com/wd/hub"
    driver = Remote(
        command_executor=remote_url,
        options=options
    )
    yield driver
    driver.quit()

BASE_URL = "http://localhost:1323/"

def test_title_exact(driver):
    driver.get(BASE_URL)
    assert driver.title == "Example Domain"
    assert "Example Domain" in driver.title
    assert len(driver.title) > 0

def test_heading_text(driver):
    driver.get(BASE_URL)
    h1 = driver.find_element(By.TAG_NAME, "h1")
    assert h1.is_displayed()
    assert h1.text == "Example Domain"
    assert "Example" in h1.text

def test_paragraph_text(driver):
    driver.get(BASE_URL)
    p = driver.find_element(By.TAG_NAME, "p")
    assert p.is_displayed()
    assert "illustrative examples" in p.text
    assert p.text.startswith("This domain is");

def test_link_href(driver):
    driver.get(BASE_URL)
    link = driver.find_element(By.CSS_SELECTOR, "a")
    href = link.get_attribute("href")
    assert href == "https://www.iana.org/domains/example?flag=.example"
    assert href.startswith("https://")
    assert ".example" in href

def test_link_text(driver):
    driver.get(BASE_URL)
    link = driver.find_element(By.CSS_SELECTOR, "a")
    assert link.text == "More information..."
    assert "More information" in link.text
    assert link.is_enabled()

def test_link_navigation(driver):
    driver.get(BASE_URL)
    link = driver.find_element(By.CSS_SELECTOR, "a")
    link.click()
    assert "iana.org" in driver.current_url
    assert driver.title != ""
    assert driver.current_url.startswith("https://")

def test_html_lang_attribute(driver):
    driver.get(BASE_URL)
    html = driver.find_element(By.TAG_NAME, "html")
    lang = html.get_attribute("lang")
    assert lang == "en"
    assert isinstance(lang, str)
    assert len(lang) == 2

def test_body_tag_present(driver):
    driver.get(BASE_URL)
    body = driver.find_element(By.TAG_NAME, "body")
    assert body is not None
    assert body.tag_name == "body"
    assert body.text != ""

def test_page_source_contains_doctype(driver):
    driver.get(BASE_URL)
    source = driver.page_source.lower()
    assert "<!doctype html>" in source
    assert "<html" in source
    assert "example domain" in source

def test_no_login_form(driver):
    driver.get(BASE_URL)
    forms = driver.find_elements(By.TAG_NAME, "form")
    assert len(forms) == 0
    assert forms == []

def test_multiple_selector_count(driver):
    driver.get(BASE_URL)
    tags = driver.find_elements(By.CSS_SELECTOR, "h1, p, a")
    assert len(tags) >= 3
    assert any(tag.tag_name == "h1" for tag in tags)
    assert any(tag.tag_name == "p" for tag in tags)

def test_parent_child_relationship(driver):
    driver.get(BASE_URL)
    parent = driver.find_element(By.TAG_NAME, "div")
    child = parent.find_element(By.TAG_NAME, "a")
    assert child in parent.find_elements(By.TAG_NAME, "a")
    assert parent.tag_name == "div"

def test_displayed_and_enabled(driver):
    driver.get(BASE_URL)
    elements = driver.find_elements(By.TAG_NAME, "a")
    for el in elements:
        assert el.is_displayed()
        assert el.is_enabled()
        assert el.get_attribute("href")

def test_css_property_background(driver):
    driver.get(BASE_URL)
    body = driver.find_element(By.TAG_NAME, "body")
    bg = body.value_of_css_property("background");
    assert bg is not None
    assert "rgb(255, 0, 0)" in bg

def test_window_handle_unique(driver):
    driver.get(BASE_URL)
    handle = driver.current_window_handle
    handles = driver.window_handles
    assert isinstance(handles, list)
    assert handle in handles

def test_execute_script(driver):
    driver.get(BASE_URL)
    title = driver.execute_script("return document.title;")
    assert title == driver.title
    assert isinstance(title, str)

def test_element_attributes(driver):
    driver.get(BASE_URL)
    h1 = driver.find_element(By.TAG_NAME, "h1")
    assert h1.get_attribute("id") == ""
    assert h1.get_attribute("class") == ""
    assert h1.get_attribute("hidden") is None

def test_timeout_implicitly_wait(driver):
    driver.implicitly_wait(10)
    driver.get(BASE_URL)
    h1 = driver.find_element(By.TAG_NAME, "h1")
    assert h1.text == "Example Domain"

def test_navigation_back_and_forward(driver):
    driver.get(BASE_URL)
    link = driver.find_element(By.CSS_SELECTOR, "a")
    link.click()
    forward_url = driver.current_url
    driver.back()
    assert driver.current_url == BASE_URL
    driver.forward()
    assert driver.current_url == forward_url

def test_screenshot_saved(driver, tmp_path):
    driver.get(BASE_URL)
    screenshot = tmp_path / "screenshot.png"
    driver.save_screenshot(str(screenshot))
    assert screenshot.exists()
    assert screenshot.stat().st_size > 0
