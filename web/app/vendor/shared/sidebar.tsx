'use client';
import { Button } from "@/components/ui/button";
import { useAuth } from "@/context/AuthContext";
import { logOut } from "@/lib/api";
import clsx from "clsx";
import { LogOut, Home, Utensils, Wallet, Store, ShoppingCart } from 'lucide-react';
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";

const links = [
  { href: '/vendor/dashboard', label: 'Dashboard', icon: Home },
  { href: '/vendor/profile', label: 'My Restaurant', icon: Store },
  { href: '/vendor/menu', label: 'Menu Builder', icon: Utensils },
  { href: '/vendor/orders', label: 'Orders', icon: ShoppingCart },
  { href: '/vendor/earnings', label: 'Earnings', icon: Wallet },
];

const Sidebar = () => {
  const { user, setUser } = useAuth();
  const pathname = usePathname();
  const router = useRouter();

  const handleLogout = async () => {
    try {
      await logOut()
      setUser(null);
    } catch (err: any) {
      console.log("Error: ", err)
    }
    router.replace('/auth/login');
  };

  return (
    <aside className="w-64 bg-white border-r border-gray-200 shadow-sm sticky top-0 self-start">
      <nav className="flex-1 p-4 space-y-1">
        {links.map(({ href, label, icon: Icon }) => (
          <Link key={href} href={href}>
            <div
              className={clsx(
                'flex items-center space-x-3 p-2 rounded-lg hover:bg-gray-100 transition-colors cursor-pointer',
                pathname === href && 'bg-gray-100 font-medium text-gray-900'
              )}
            >
              <Icon className="h-5 w-5 text-gray-600" />
              <span>{label}</span>
            </div>
          </Link>
        ))}
      </nav>

      <div className="p-4 border-t border-gray-100">
        <Button
          variant="destructive"
          onClick={handleLogout}
          className="w-full flex items-center gap-2 justify-center"
        >
          <LogOut className="h-4 w-4" />
          Sign Out
        </Button>
      </div>
    </aside>
  );
};

export default Sidebar;

