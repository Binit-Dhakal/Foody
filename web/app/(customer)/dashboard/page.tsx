import ProtectedRoute from '@/components/ProtectedRoutes';
import UserDetail from './user';

export default function CustomerDashboard() {
    return (
        <ProtectedRoute allowedRoles={['customer']}>
            <div>
                <UserDetail />
            </div>
        </ProtectedRoute>
    );
}

