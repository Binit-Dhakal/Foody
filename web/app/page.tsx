"use client";

import { useAuth } from "@/context/AuthContext";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function Home() {
  const { user, loading } = useAuth();
  const router = useRouter();

  useEffect(() => {
    if (loading) return; // Wait for session check

    if (!user) {
      router.replace("/auth/login");
      return;
    }

    if (user.role === "customer") {
      router.replace("/dashboard");
    } else if (user.role === "vendor") {
      router.replace("/vendor/dashboard");
    }
  }, [user, loading, router]);

  return (
    <div className="flex items-center justify-center min-h-screen">
      <p>Redirecting...</p>
    </div>
  );
}

