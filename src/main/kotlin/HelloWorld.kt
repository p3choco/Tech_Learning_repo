fun main() {
    println("Hello from Kotlin!")
    println("Python is also available in this container.")
    println("Checking SQLite JDBC driver:")

    try {
        Class.forName("org.sqlite.JDBC")
        println("SQLite JDBC Driver is loaded!")
    } catch (ex: ClassNotFoundException) {
        println("Could not load SQLite JDBC Driver: ${ex.message}")
    }
}
