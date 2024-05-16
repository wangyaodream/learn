import profilePic from './assets/profile.jpeg';

function Card() {

    return (
        <div className="card">
            <img className="card-image" src={profilePic} alt="profile picture" />
            <h2 className='card-title'>Wang Yao</h2>
            <p className='card-description'>A developer</p>
        </div>
    );
}

export default Card;