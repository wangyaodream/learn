function greeter(person) {
    return "Hello Mr(Miss) " + person.firstName;
}
var user = {
    firstName: "wang",
    LastName: "yao"
};
console.log(greeter(user));
