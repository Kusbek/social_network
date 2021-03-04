import {
    ref
} from "vue"

const error = ref(null)

const signup = async (username, email, firstName, lastName, birthDate, aboutMe, password) => {
    error.value = null
    let user = {
        username: username,
        email: email,
        first_name: firstName,
        last_name: lastName,
        birth_date: birthDate,
        about_me: aboutMe,
        password: password,
    }
    console.log(JSON.stringify(user))
    // error.value = "Failed to signup, please try later"
}

const useSignup = () => {
    return {
        error,
        signup
    }
}

export default useSignup