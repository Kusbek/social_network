import {
    ref
} from "vue"

const error = ref(null)

const signup = async (username, email, firstName, lastName, birthDate, aboutMe, password) => {
    error.value = null
    let body = {
        username: username,
        email: email,
        first_name: firstName,
        last_name: lastName,
        birth_date: birthDate,
        about_me: aboutMe,
        password: password,
    }
    try {
        let res = await fetch("http://localhost:8080/signup", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(body)
        })
        if (!res.ok) {
            throw Error("Failed to sign up")
        }
    } catch (err) {
        error.value = err.message
        console.log(err.message)
    }
}

const useSignup = () => {
    return {
        error,
        signup
    }
}

export default useSignup