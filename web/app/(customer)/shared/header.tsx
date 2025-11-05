'use client';
import { useAuth } from "@/context/AuthContext";
import Image from "next/image";

export default function CustomerHeader() {
  const { user } = useAuth();

  return (
    <header className="bg-white border-b border-gray-200 px-6 py-4 flex items-center justify-between">
      <div>
        <h1 className="text-xl font-semibold text-gray-800">
          👋 Hi {user?.name || 'there'}!
        </h1>
        <p className="text-gray-500 text-sm">
          Ready to explore your favorite meals today?
        </p>
      </div>

      <div className="flex items-center gap-4">
        <Image
          src="/default-profile.PNG"
          alt="Profile"
          width={40}
          height={40}
          className="rounded-full border border-gray-300"
        />
      </div>
    </header>
  );
}

