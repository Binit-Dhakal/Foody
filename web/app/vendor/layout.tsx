import ProtectedRoute from '@/components/ProtectedRoutes';
import Sidebar from './shared/sidebar';
import Header from './shared/header';
import { VendorProvider } from '@/context/VendorContext';

export default function VendorLayout({ children }: { children: React.ReactNode }) {
  return (
    <VendorProvider>
      <ProtectedRoute allowedRoles={['vendor']}>
        <div className="min-h-screen flex flex-col bg-gray-50">
          <div className="w-full">
            <Header />
          </div>

          <div className="w-full h-10" />

          <div className="flex flex-1 content-center">
            <div className="flex w-full px-40 gap-8">
              <Sidebar />
              <main className="flex-1 overflow-y-auto bg-gray-50 rounded-xl shadow-sm">
                {children}
              </main>
            </div>
          </div>
        </div>

      </ProtectedRoute>
    </VendorProvider>
  );
}

