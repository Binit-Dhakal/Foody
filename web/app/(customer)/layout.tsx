import ProtectedRoute from '@/components/ProtectedRoutes';
import Sidebar from './shared/sidebar';
import CustomerHeader from './shared/header';

export default function VendorLayout({ children }: { children: React.ReactNode }) {
  return (
    <ProtectedRoute allowedRoles={['customer']}>
      <div className="min-h-screen flex bg-gray-50">
        <Sidebar />
        <main className="flex-1 flex flex-col">
          <CustomerHeader />
          <div className="p-6 flex-1 overflow-y-auto">{children}</div>
        </main>
      </div>
    </ProtectedRoute>
  );
}

