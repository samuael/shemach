import React from 'react'

const ErrorAlter = (props) => {
    return (
        <div className="error-log">
            <button onClick={props.clearError}>{props.errorMessage} <i className="error"></i></button>
        </div>
    )
}

export default ErrorAlter

