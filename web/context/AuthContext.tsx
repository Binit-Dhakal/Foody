"use client";

import { createContext, useContext, useState, useEffect, ReactNode } from "react";
import { fetchSession } from "@/lib/api";
import { useRouter } from "next/navigation";

interface User {
  email: string;
  role: "admin" | "vendor" | "customer";
  name: string;
}

interface AuthContextType {
  user: User | null;
  loading: boolean;
  setUser: (user: User | null) => void;
  reloadSession: () => Promise<void>; // renamed to be semantically clear
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  const reloadSession = async () => {
    try {
      const res = await fetchSession();
      setUser(res.data);
    } catch {
      setUser(null);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    reloadSession();
  }, []);

  return (
    <AuthContext.Provider value={{ user, loading, setUser, reloadSession }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth must be used within an AuthProvider");
  return context;
}

export function useRoleRedirect() {
  const { user } = useAuth();
  const router = useRouter();

  return () => {
    if (!user) router.replace('/auth/login');
    else if (user.role === 'customer') router.replace('/dashboard');
    else if (user.role === 'vendor') router.replace('/vendor/dashboard');
    else if (user.role === 'admin') router.replace('/admin/dashboard');
  };
}
