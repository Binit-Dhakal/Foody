"use client";

import { ReactNode, useEffect } from 'react';
import { useAuth } from '@/context/AuthContext';
import { useRouter } from 'next/navigation';

interface ProtectedRouteProps {
    allowedRoles: ('admin' | 'vendor' | 'customer')[];
    children: ReactNode;
}

export default function ProtectedRoute({ allowedRoles, children }: ProtectedRouteProps) {
    const { user, loading } = useAuth();
    const router = useRouter();

    useEffect(() => {
        if (!loading) {
            if (!user) {
                // Not logged in → redirect to login
                router.replace('/auth/login');
            } else if (!allowedRoles.includes(user.role)) {
                // Logged in but not allowed → redirect to their dashboard/home
                switch (user.role) {
                    case 'admin':
                        router.replace('/admin');
                        break;
                    case 'vendor':
                        router.replace('/vendor/dashboard');
                        break;
                    case 'customer':
                        router.replace('/');
                        break;
                }
            }
        }
    }, [user, loading, router, allowedRoles]);

    if (loading || !user || !allowedRoles.includes(user.role)) {
        // Optional: render loading spinner while checking
        return <div>Loading...</div>;
    }

    return <>{children}</>;
}

