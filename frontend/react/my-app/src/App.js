import * as React from 'react'


// const List = ({ list }) => (
//     <ul>
//         {list.map((item) => ( 
//             <Item key={item.objectID} item={item} />
//         ))}
//     </ul>
// );
const List = ({ list }) => (
    // <ul>
    //     {list.map((item) => ( 
    //         <Item 
    //             key={item.objectID} 
    //             title={item.title}
    //             url={item.url}
    //             author={item.author}
    //             num_comments={item.num_comments}
    //             points={item.points} />
    //     ))}
    // </ul>
    <ul>
        {list.map(({ objectID, ...item }) => (
            <Item key={objectID} {...item} />
        ))}
    </ul>
);

// const Item = ({ item }) => ( 
//     <li>
//         <span>
//            <a href={item.url}>{item.title}</a>
//         </span>
//         <span>{item.author}</span>
//         <span>{item.num_comments}</span>
//         <span>{item.points}</span>
//     </li>
// );
const Item = ({ title, url, author, num_comments, points }) => (
    <li>
        <span>
           <a href={url}>{title}</a>
        </span>
        <span>{author}</span>
        <span>{num_comments}</span>
        <span>{points}</span>
    </li>
);



const Search = ({ search, onSearch }) => {
    // const [] = React.useState('');
    //
    // const handleChange = (event) => {
    //      setSearchTerm(event.target.value);
    // };

    // const {search, onSearch} = props;

    return (
        <div>
            <label htmlFor='search'>Search: </label>
            <input 
                id='search' 
                type='text' 
                value={search}
                onChange={onSearch}/>
        </div>
    )
}

const App = () => {
    const stories = [
        {
            title: 'React',
            url: 'https://reactjs.org/',
            author: 'Jordan Walke',
            num_comments: 3,
            points: 4,
            objectID: 0,
        },
        {
            title: "Redux",
            url: "https://redux.js.org/",
            author: "Dan Abramov, Andrew Clark",
            num_comments: 2,
            points: 5,
            objectID:1,
        },
    ];

    const [searchTerm, setSearchTerm] = React.useState("");

    const handleSearch = (event) => {
        setSearchTerm(event.target.value);
    }

    const searchedStories = stories.filter(function (story) {
        return story.title.toLowerCase().includes(searchTerm.toLowerCase());
    })

    return (
        <div>
            <h1>My Hacker Stories</h1>

            <Search search={searchTerm} onSearch={handleSearch}/>

            <hr />
            <List list={searchedStories}/>
        </div>
    )
}
export default App;
