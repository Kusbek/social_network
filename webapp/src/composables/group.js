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
            group.value = {
                ...data
            }
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }

    const createGroup = async (title, description) => {
        let body = {
            title: title,
            description: description
        }
        error.value = null
        try {
            let res = await fetch('/api/group', {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(body)
            })
            if (!res.ok) {
                throw Error("Could not create group")
            }
            let newGroup = await res.json()

            return newGroup
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }

    const groups = ref([])
    const getGroups = async () => {
        error.value = null
        try {
            let res = await fetch(`/api/groups`)
            if (!res.ok) {
                throw Error("Could not fetch groups")
            }
            let data = await res.json()
            groups.value = data.group_list
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }

    return {
        error,
        group,
        groups,
        getGroup,
        createGroup,
        getGroups
    }
}


export default useGroup