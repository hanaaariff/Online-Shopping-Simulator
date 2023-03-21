package main

import (
  "os" // to open text file
  "bufio" // to read data from text file
  "strings"
  "strconv"
  "fmt"
)

// Struct to hold descriptors of a product
type product struct {
  id int
  name string
  brand string
  quantity int
  price string
}

// Slices to hold copy of grocery store and shopping cart
var store []product
var cart []product

func print_contents(s []product) {
  for i := 0; i < len(s); i++ {
    fmt.Printf("ID: %d, Item: %s, Brand: %s, Quantity: %d, Price: $%s\n", s[i].id, s[i].name, s[i].brand, s[i].quantity, s[i].price)
  }
}

func add_item_to_cart() {
  fmt.Println("\nGrocery Store: ")
  print_contents(store)
  var answer int
  fmt.Print("\nEnter the ID number of the item you wish to add to the shopping cart: ")
  fmt.Scan(&answer)

  is_in_cart := false
  for i := 0; i < len(cart); i++ {
    if answer == cart[i].id {
      is_in_cart = true
      cart[i].quantity++
      store[answer].quantity--
    }
  }
  if is_in_cart == false {
    add_p := product {
      id: store[answer].id,
      name: store[answer].name,
      brand: store[answer].brand,
      quantity: 1,
      price: store[answer].price,
    }
    store[answer].quantity--
    cart = append(cart, add_p)
  }
}

func remove_item_from_cart() {
  fmt.Println("\nShopping Cart: ")
  print_contents(cart)
  var answer int
  fmt.Print("\nEnter the ID number of the item you wish to remove from the shopping cart: ")
  fmt.Scan(&answer)

  for i := 0; i < len(cart); i++ {
    if cart[i].id == answer {
      if cart[i].quantity > 0 {
        cart[i].quantity--
        store[answer].quantity++
      }
    }
  }
}

func home_menu() {
  fmt.Println("\nMenu:")
  fmt.Println("(1) View shopping cart")
  fmt.Println("(2) Add item to cart")
  fmt.Println("(3) Remove item from cart")
  fmt.Println("(4) Check out and pay")
  var answer int
  fmt.Print("\nEnter the number of the action you wish to take: ")
  fmt.Scan(&answer)
  if answer == 1 {
      print_contents(cart)
      home_menu()
  } else if answer == 2 {
      add_item_to_cart()
      home_menu()
  } else if answer == 3 {
      remove_item_from_cart()
      home_menu()
  } else if answer == 4 {
      fmt.Println("\nThanks for shopping with us!")
  } else {
      fmt.Println("\nError: Please enter one of the listed number options.")
      home_menu()
  }
}

func main() {
  // Open grocerystore.txt
  grocerystore, ferr := os.Open("grocerystore.txt")
  // Error checking
  if ferr != nil {
    panic(ferr)
  }

  // Create buffer to read the data into and load file into buffer
  scanner := bufio.NewScanner(grocerystore)

  counter := 0

  // Read data
  for scanner.Scan() {
    line := scanner.Text()
    descriptors := strings.Split(line, ",")

    quantity_int, err := strconv.Atoi(descriptors[2])
    if err == nil {
      p := product {
        id: counter,
        name: descriptors[0],
        brand: descriptors[1],
        quantity: quantity_int,
        price: descriptors[3],
      }
      store = append(store, p)
      counter = counter + 1
    }
  }

  // Welcome the customer
  fmt.Println("\nWelcome to the Grocery Store!\n")
  home_menu()

  // Create update grocery store file
  new_file, ferr2 := os.Create("updated_grocerystore.txt")
  if ferr2 != nil {
    panic(ferr)
  }
  for i := 0; i < len(store); i++ {
    quantity_string := strconv.Itoa(store[i].quantity)
    str := store[i].name+","+store[i].brand+","+quantity_string+","+store[i].price+"\n"
    new_file.WriteString(str)
  }
  ferr3 := grocerystore.Close()
  if ferr3 != nil {
    panic(ferr3)
  }
  ferr4 := new_file.Close()
  if ferr4 != nil {
    panic(ferr4)
  }
}
