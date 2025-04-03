package controllers

import models.Category
import play.api.mvc.*
import play.api.libs.json.*

import javax.inject.*
import scala.collection.mutable.ListBuffer

@Singleton
class CategoriesController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  implicit val categoryFormat: OFormat[Category] = Json.format[Category]

  private val categories = ListBuffer(
    Category(1, "Kategoria jax"),
    Category(2, "Kategoria max")
  )

  def listCategories: Action[AnyContent] = Action {
    Ok(Json.toJson(categories))
  }

  def getCategory(id: Int): Action[AnyContent] = Action {
    categories.find(_.id == id) match {
      case Some(category) => Ok(Json.toJson(category))
      case None => NotFound(s"Nie znaleziono kategorii o id $id")
    }
  }

  def addCategory: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest,
      category => {
        categories += category
        Created(Json.toJson(category))
      }
    )
  }

  def updateCategory(id: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest,
      updatedCategory => {
        categories.indexWhere(_.id == id) match {
          case -1 => NotFound(s"Nie znaleziono kategorii o id $id")
          case index =>
            categories.update(index, updatedCategory)
            Ok(Json.toJson(updatedCategory))
        }
      }
    )
  }

  def deleteCategory(id: Int): Action[AnyContent] = Action {
    categories.indexWhere(_.id == id) match {
      case -1 => NotFound(s"Nie znaleziono kategorii o id $id")
      case index =>
        categories.remove(index)
        NoContent
    }
  }
}