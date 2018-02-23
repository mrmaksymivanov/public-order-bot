package telegram

import (
  "github.com/Syfaro/telegram-bot-api"
  "public-order-bot/dao"
)

func getResponse(update *tgbotapi.Update) {
  switch update.Message.Command() {
    case "create": createOrder(update)
    case "list": ordersList(update)
    case "add": addItem(update)
    case "remove": removeItem(update)
    case "result": orderResult(update)
    case "close": closeOrder(update)
  }
}

func createOrder(update *tgbotapi.Update) {
  orderName := update.Message.CommandArguments()
  var response string
  if orderName == "" {
    response = getResponseText("create_fail_order_name", nil)
  } else {
    order, err := dao.CreateOrder(string(update.Message.Chat.ID), string(update.Message.From.UserName), orderName)
    check(err)
    responseData := templateData {
      "orderName": order.Name,
      "orderId": order.Id,
    }
    response = getResponseText("create_success", responseData)
  }
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func ordersList(update *tgbotapi.Update) {
  orders, err := dao.ListOrders(string(update.Message.Chat.ID))
  check(err)
  responseData := templateData {
    "orderList": orders,
  }
  response := getResponseText("orders_list", responseData)
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  telegramBot.Send(msg)
}

func addItem(update *tgbotapi.Update) {
  response := "Item added to order"
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func removeItem(update *tgbotapi.Update) {
  response := "Item has been removed"
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func orderResult(update *tgbotapi.Update) {
  response := "Order result"
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func closeOrder(update *tgbotapi.Update) {
  response := "Order has been closed"
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}
