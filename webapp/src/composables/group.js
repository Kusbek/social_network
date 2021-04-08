import {
    ref
} from '@vue/reactivity'

const useGroup = () => {
    const error = ref(null)


    const group = ref(null)

    const getGroup = async (id) => {
        // console.log(window.location.host + `/group?group_id=${id}`)
        error.value = null
        try {
            let res = await fetch(`/api/group?group_id=${id}`)
            if (!res.ok) {
                throw Error("Could not fetch group")
            }
            let data = await res.json()
            console.log(data)
            group.value = {
                ...data
            }
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }

    return {
        error,
        group,
        getGroup,
    }
}


export default useGroup