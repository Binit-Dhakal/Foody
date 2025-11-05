import VendorProfileForm from "./profileForm";

export default function VendorProfilePage() {
  const vendor = {
    name: "Pizza Palace",
    address: "Dhapakhel, Lalitpur",
    country: "Nepal",
    state: "Bagmati",
    city: "Lalitpur",
    pinCode: "44700",
    logo: "/default-profile.PNG",
    coverPhoto: "/default-cover.PNG",
  };

  return (
    <VendorProfileForm vendor={vendor} />
  );
}

