package controllers

import models.CartItem
import play.api.mvc.*
import play.api.libs.json.*

import javax.inject.*
import scala.collection.mutable.ListBuffer

@Singleton
class CartController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  implicit val cartItemFormat: OFormat[CartItem] = Json.format[CartItem]

  private val cart = ListBuffer(
    CartItem(1324, 123),
    CartItem(9826, 321)
  )

  def listCartItems: Action[AnyContent] = Action {
    Ok(Json.toJson(cart))
  }

  def addCartItem: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[CartItem].fold(
      errors => BadRequest,
      cartItem => {
        cart += cartItem
        Created(Json.toJson(cartItem))
      }
    )
  }

  def updateCartItem(productId: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[CartItem].fold(
      errors => BadRequest,
      updatedCartItem => {
        cart.indexWhere(_.productId == productId) match {
          case -1 => NotFound(s"Nie znaleziono produktu o id $productId w koszyku")
          case index =>
            cart.update(index, updatedCartItem)
            Ok(Json.toJson(updatedCartItem))
        }
      }
    )
  }

  def deleteCartItem(productId: Int): Action[AnyContent] = Action {
    cart.indexWhere(_.productId == productId) match {
      case -1 => NotFound(s"Nie znaleziono produktu o id $productId w koszyku")
      case index =>
        cart.remove(index)
        NoContent
    }
  }
}