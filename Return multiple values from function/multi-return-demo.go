package main;

import "fmt"

func GetDomainDetails() (string, string) {
    return "turreta.com", "projecting knowledge"
}

func main() {
    domain, title := GetDomainDetails()
    fmt.Println(domain)
    fmt.Println(title)
}


