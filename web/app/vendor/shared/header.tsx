'use client';
import { useVendor } from "@/context/VendorContext";
import Image from "next/image";

const VendorHeader = () => {
  const { vendor } = useVendor()
  console.log(vendor)

  return (
    <div className="relative w-full h-64 md:h-80 lg:h-96">
      <Image
        src={vendor?.coverPhoto || '/default-cover.PNG'}
        alt="Cover"
        fill
        className="object-cover"
        priority
      />

      <div className="absolute inset-0 bg-gradient-to-t from-black/70 to-transparent" />

      <div className="absolute bottom-20 left-40 flex items-center space-x-4 text-white">
        <div className="relative w-40 h-40 rounded-full border-4 border-white overflow-hidden">
          <Image
            src={vendor?.logo || '/default-profile.PNG'}
            alt="Logo"
            fill
            className="object-cover"
          />
        </div>
        <div>
          <h1 className="text-4xl font-bold">{vendor?.name}</h1>
          <p className="text-md text-gray-200">{vendor?.location}</p>
        </div>
      </div>
    </div>)
}

export default VendorHeader;
