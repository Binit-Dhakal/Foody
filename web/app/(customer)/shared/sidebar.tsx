'use client';

import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import clsx from "clsx";
import { Button } from "@/components/ui/button";
import { Home, ShoppingBag, Heart, MapPin, Settings, LogOut } from "lucide-react";
import { useAuth } from "@/context/AuthContext";
import { logOut } from "@/lib/api";

const links = [
  { href: '/dashboard', label: 'Dashboard', icon: Home },
  { href: '/dashboard/orders', label: 'My Orders', icon: ShoppingBag },
  { href: '/dashboard/favorites', label: 'Favorites', icon: Heart },
  { href: '/dashboard/addresses', label: 'Addresses', icon: MapPin },
  { href: '/dashboard/settings', label: 'Settings', icon: Settings },
];

const Sidebar = () => {
  const pathname = usePathname();
  const router = useRouter();
  const { setUser } = useAuth();

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
    <aside className="w-64 bg-white border-r border-gray-200 flex flex-col shadow-sm">
      <div className="p-6 text-2xl font-semibold text-gray-800 border-b border-gray-100">
        Foody
      </div>

      <nav className="flex-1 p-4 space-y-1">
        {links.map(({ href, label, icon: Icon }) => (
          <Link key={href} href={href}>
            <div
              className={clsx(
                'flex items-center space-x-3 p-2 rounded-lg hover:bg-gray-100 transition-colors cursor-pointer',
                pathname === href && 'bg-gray-100 text-gray-900 font-medium'
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

