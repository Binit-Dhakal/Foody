import ProtectedRoute from '@/components/ProtectedRoutes';

export default function VendorDashboard() {
    return (
        <ProtectedRoute allowedRoles={['vendor']}>
            <div>
                <h1>Vendor Dashboard</h1>
            </div>
        </ProtectedRoute>
    );
}

