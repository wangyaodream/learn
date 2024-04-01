import * as React from 'react'
import { useState } from 'react';
// const welcome = {
//     greeting: 'Hey',
//     title: 'React'
// }

function Square({value, onSquareClick}) {
    // const [value, setValue] = useState(null);
    // function handleClick() {
    //     // console.log("clicked!")
    //     // if (value === "X") {
    //     //     setValue('')
    //     // } else {
    //     //     setValue('X');
    //     // }
    //     setValue('X')
    // }

    // return (
    //     <button
    //     className='square'
    //     onClick={handleClick}
    //     >
    //     {value}
    //     </button>
    // )
    return (
        <button 
            className='square' 
            onClick={onSquareClick}>
                {value}
        </button>);
}

function App() {
    const [squares, setSquares] = useState(Array(9).fill(null));

    function handleClick(i) {
        const nextSquares = squares.slice();
        nextSquares[i] = 'X';
        setSquares(nextSquares);
    }
    return (
        <>
            <div className='board-row'>
                <Square value={squares[0]} onSquareClick={() => handleClick(0)}/>
                <Square value={squares[1]} onSquareClick={() => handleClick(1)}/>
                <Square value={squares[2]} onSquareClick={() => handleClick(2)}/>
            </div>
            <div className='board-row'>
                <Square value={squares[3]} onSquareClick={() => handleClick(3)}/>
                <Square value={squares[4]} onSquareClick={() => handleClick(4)}/>
                <Square value={squares[5]} onSquareClick={() => handleClick(5)}/>
            </div>
            <div className='board-row'>
                <Square value={squares[6]} onSquareClick={() => handleClick(6)}/>
                <Square value={squares[7]} onSquareClick={() => handleClick(7)}/>
                <Square value={squares[8]} onSquareClick={() => handleClick(8)}/>
            </div>
        </>
    ); 
}

export default App;
