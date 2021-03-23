import {
    ref
} from 'vue'

const error = ref(null)

const validateUsername = (username) => {
    if (username.length) {
        if (username.length < 5 || username.length > 20) {
            error.value = "username should be within 5 and 20 symbols"
            return
        }
        let re = new RegExp(/^[a-zA-Z0-9]+$/, "g")
        if (!username.match(re)) {
            error.value = "username should be in latin symbols"
            return
        }
    }
}

const validateEmail = (email) => {
    let re = new RegExp(/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/, "g")
    if (!re.test(email)) {
        error.value = 'email is invalid'
    }
}


const validatePassword = (password, confirmPassword) => {
    if (password !== confirmPassword) {
        error.value = 'confirmed password is not equal to password'
        return
    }
}

const validateUser = (username, email, password, confirmPassword) => {
    error.value = null
    validatePassword(password, confirmPassword);
    validateEmail(email);
    validateUsername(username);
}

const useValidators = () => {
    return {
        error,
        validateUsername,
        validateEmail,
        validatePassword,
        validateUser
    }
}


export default useValidators