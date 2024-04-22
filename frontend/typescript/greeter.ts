function greeter(person: Person) {
    return "Hello Mr(Miss) " + person.firstName;
}

let user = {
    firstName: "wang",
    LastName: "yao"
}


interface Person {
    firstName: string;
    LastName: string;
}

console.log(greeter(user));


// Class
class Student {
    fullName: string;
    constructor(public firstName, public middleIn)
}
