import ProtectedRoute from '@/components/ProtectedRoutes';

export default function VendorLayout({ children }: { children: React.ReactNode }) {
  return (
    <ProtectedRoute allowedRoles={['customer']}>
      {children}
    </ProtectedRoute>
  );
}

