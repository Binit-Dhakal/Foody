"use client";
import { getVendorSummary } from "@/lib/vendor-api";
import { createContext, useContext, useEffect, useState, ReactNode } from "react";

interface Vendor {
  name: string;
  location: string;
  logo: string | null;
  coverPhoto: string | null;
}

interface VendorContextType {
  vendor: Vendor | null;
  loading: boolean;
  reloadVendor: () => Promise<void>;
  setVendor: (v: Vendor | null) => void;
}

const VendorContext = createContext<VendorContextType | undefined>(undefined);

export function VendorProvider({ children }: { children: ReactNode }) {
  const [vendor, setVendor] = useState<Vendor | null>(null);
  const [loading, setLoading] = useState(true);

  const reloadVendor = async () => {
    try {
      const res = await getVendorSummary()
      const data = await res.data;

      setVendor({
        name: data.name,
        location: data.location,
        logo: data.logo,
        coverPhoto: data.cover_photo,
      });
    } catch (err) {
      const vendor = {
        name: "Pizza Place",
        location: "Dhapakhel, Lalitpur",
        logo: null,
        coverPhoto: null
      }
      setVendor(vendor)
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    reloadVendor();
  }, []);

  return (
    <VendorContext.Provider value={{ vendor, loading, reloadVendor, setVendor }}>
      {children}
    </VendorContext.Provider>
  );
}

export function useVendor() {
  const context = useContext(VendorContext);
  if (!context) throw new Error("useVendor must be used inside VendorProvider");
  return context;
}

