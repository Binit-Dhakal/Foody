'use client';
import { useAuth } from "@/context/AuthContext"

const UserDetail = () => {
    const { user } = useAuth()
    return (
        <>
            Welcome {user?.name} - Role-{user?.email}
        </>
    )
}

export default UserDetail;
