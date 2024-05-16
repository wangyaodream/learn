import Header from "./Header"
import Footer from "./Footer"
import Card from "./Card"
import Button from "./Button/Button"
import Student from "./Student"

function App() {
    return (
        <div>
            <Header />
            <Card></Card>
            <Button></Button>
            <Student name="yao" age="20" isStudent={false}></Student>
            <Student name="he" age={30} isStudent={true}></Student>
            <Footer></Footer> 
        </div>
    );
}

export default App
