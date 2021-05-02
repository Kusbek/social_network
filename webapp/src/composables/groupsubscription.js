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
            let res = await fetch(`/api/group/invite?group_id=${groupId}`, {
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
                if (res.status == 418) {
                    throw Error("you can't invite owner of the group")
                }
                if (res.status == 403) {
                    throw Error("you are not allowed to invite")
                }
                throw Error("Something went wrong")
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
            let res = await fetch(`/api/group/invite/accept`, {
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

    const groupMemberList = ref([])
    const getGroupMemberList = async(groupId) => {
        error.value = null
        try {
            let res = await fetch(`/api/group/members?group_id=${groupId}`)
            if (!res.ok) {
                throw Error("Failed to get group members")
            }
            let data = await res.json()
            groupMemberList.value = data.group_members
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }    
    }

    const isGroupMember = ref(false)
    const reqIsPending = ref(false)
    const checkIfGroupMember = async(userId, groupId) => {
        error.value = null
        
        try {
            let res = await fetch(`/api/group/ismember?group_id=${groupId}&user_id=${userId}`)
            if (!res.ok) {
                throw Error("Failed to check if group member")
            }
            let data = await res.json()
            isGroupMember.value = data.is_group_member
            reqIsPending.value = data.request_is_pending
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }  
    }


    const requestToJoin = async(groupId) => {
        error.value = null
        let body = {
            group_id: Number(groupId),
        }
        try {
            let res = await fetch(`/api/group/invite/join`, {
                method: "POST",
                body: JSON.stringify(body)
            })
            if (!res.ok) {
                throw Error("Failed to send a join request")
            }

        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }  
    }

    return {
        error,
        groupInviteList,
        groupMemberList,
        isGroupMember,
        reqIsPending,
        requestToJoin,
        checkIfGroupMember,
        invite,
        getGroupInviteList,
        acceptInvite,
        getGroupMemberList
    }
}
export default useGroupSubscription