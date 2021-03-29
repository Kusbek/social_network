import {
    ref
} from "@vue/reactivity"
const profile = ref(null)
const error = ref(null)

const load = async (id) => {
    error.value = null
    try {
        let res = await fetch(`./api/user?id=${id}`)
        let data = await res.json()
        if (!res.ok) {
            throw Error("Could not fetch profile")
        }
        profile.value = {
            ...data
        }
    } catch (e) {
        console.log(e.message)
        error.value = e.message
    }
    // profile.value = {
    //     id: id,
    //     username: "kusbek_test",
    //     email: "kusbek1994_test@gmail.com",
    //     firstName: "Bekarys_test",
    //     lastName: "Kuspan_test",
    //     birthDate: "1994-09-18",
    //     aboutMe: "Test testinson motherfucking testicles for mistestress",
    //     pathToPhoto: "/img/images/avatars/default.jpg"
    // }
}

const setPublicity = async (isPublic) => {
    error.value = null
    let body = {
        is_public: isPublic,
    }
    try {
        let res = await fetch(`./api/user/setprofilevisibility`, {
            method: "PATCH",
            body: JSON.stringify(body)
        })
        if (!res.ok) {
            throw Error("Could not update publicity")
        }
    } catch (e) {
        console.log(e.message)
        error.value = e.message
    }
}




const useProfile = () => {
    return {
        profile,
        error,
        setPublicity,
        load,
    }
}

export default useProfile