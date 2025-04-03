package controllers

import models.Product
import play.api.mvc.*
import play.api.libs.json.*

import javax.inject.*
import scala.collection.mutable.ListBuffer

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  private var products: List[Product] = List(
    Product(1, "Produkt A", 10.0),
    Product(2, "Produkt B", 20.0)
  )

}