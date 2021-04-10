import {
    ref
} from "@vue/reactivity"

const error = ref(null)
const upload = async (file) => {
    error.value = null
    if (file) {
        const formData = new FormData()
        formData.append('myFile', file)
        try {
            let res = await fetch('/api/upload', {
                method: "POST",
                body: formData
            })
            let data = await res.json()
            return data.path_to_photo
        } catch (e) {
            error.value = e.message
        }
    }
}
const useFile = () => {
    return {
        error,
        upload
    }
}

export default useFile