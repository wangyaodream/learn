import PropTypes from 'prop-types';

function Student(props) {
    return(
        <div className="student">
            <p>Name: {props.name}</p>
            <p>Age: {props.age}</p>
            <p>Student: {props.isStudent ? "yes" : "no"}</p>
        </div>
    );
}

// 规定传入的参数类型
Student.propTypes = {
    name: PropTypes.string,
    age: PropTypes.number,
    isStudent: PropTypes.bool
}

// 设置默认值
Student.defaultProps = {
    name: "Anonymous",
    age: 0,
    isStudent: false
}

export default Student;