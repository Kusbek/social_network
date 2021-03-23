import {
    ref
} from "@vue/reactivity"
const profile = ref(null)
const error = ref(null)
const load = async (id) => {
    try{
        let res = await fetch(`./api/user?id=${id}`)
        let data = await res.json()
        profile.value = {
            ...data
        }
    }catch(e){
        console.log(e.message)
        error.value = "Could not fetch profile"
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

const useProfile = () => {
    return {
        profile,
        error,
        // upload,
        load,
    }
}

export default useProfile