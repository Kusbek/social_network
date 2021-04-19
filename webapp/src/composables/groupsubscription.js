import {
    ref
} from "@vue/reactivity"

const useGroupSubscription = () => {
    const error = ref(null)
    const invite = async (nickmail, groupId) => {
        error.value = null
        let body = {
            nickmail: nickmail,
            group_id: groupId
        }
        try {
            let res = await fetch('/api/group/invite', {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(body)
            })
            if (!res.ok) {
                if (res.status == 404) {
                    throw Error("could not find such user")
                } 
                if (res.status == 400) {
                    throw Error("you can't invite owner of the group")
                }
            }
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }

    const groupInviteList = ref([])
    const getGroupInviteList = async() => {
        error.value = null
        try {
            let res = await fetch(`/api/group/invites`)
            if (!res.ok) {
                throw Error("Failed to get group invites")
            }
            let data = await res.json()
            groupInviteList.value = data.group_invites
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }    
    }

    const acceptInvite = async(groupId) => {
        error.value = null
        let body = {
            group_id: groupId,
        }
        try {
            let res = await fetch(`/api/group/invite`, {
                method: "PUT",
                body: JSON.stringify(body)
            })
            if (!res.ok) {
                throw Error("Failed to accept group invite")
            }

        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }   
    }

    return {
        error,
        groupInviteList,
        invite,
        getGroupInviteList,
        acceptInvite,
    }
}
export default useGroupSubscription