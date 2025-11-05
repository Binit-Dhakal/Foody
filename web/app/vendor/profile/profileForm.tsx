'use client';
import { useState } from "react";
import { vendorSchema, VendorFormData } from "@/types/vendor";
import InputField from "@/components/ui/inputField";
import ImageUploadField from "@/components/ui/imageUploadField";
import { Button } from "@/components/ui/button";

export default function VendorProfileForm({ vendor }: { vendor: Partial<VendorFormData> }) {
  const [form, setForm] = useState<VendorFormData>({
    name: vendor.name || "",
    address: vendor.address || "",
    country: vendor.country || "",
    state: vendor.state || "",
    city: vendor.city || "",
    pinCode: vendor.pinCode || "",
    latitude: vendor.latitude || "",
    longitude: vendor.longitude || "",
    logo: null,
    coverPhoto: null,
    vendorLicense: null,
  });

  const [errors, setErrors] = useState<Record<string, string>>({});
  const [loading, setLoading] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleFileChange = (name: keyof VendorFormData, file: File | null) => {
    setForm((prev) => ({ ...prev, [name]: file }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const validation = vendorSchema.safeParse(form);
    if (!validation.success) {
      const fieldErrors: Record<string, string> = {};
      validation.error.errors.forEach((err) => {
        const field = err.path[0] as string;
        fieldErrors[field] = err.message;
      });
      setErrors(fieldErrors);
      return;
    }

    setErrors({});
    setLoading(true);

    const formData = new FormData();
    Object.entries(form).forEach(([key, value]) => {
      if (value) formData.append(key, value as any);
    });

    const res = await fetch("/api/vendor/profile", {
      method: "PUT",
      body: formData,
    });

    setLoading(false);
    if (res.ok) alert("Profile updated successfully!");
    else alert("Failed to update profile.");
  };

  const fields = [
    { label: "Restaurant Name", name: "name" },
    { label: "Address", name: "address" },
    { label: "Country", name: "country" },
    { label: "State", name: "state" },
    { label: "City", name: "city" },
    { label: "Pin Code", name: "pinCode" },
    { label: "Latitude", name: "latitude" },
    { label: "Longitude", name: "longitude" },
  ];

  return (
    <form onSubmit={handleSubmit} className="space-y-6 max-w-2xl mx-auto p-6 bg-white rounded-xl shadow">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {fields.map(({ label, name }) => (
          <InputField
            key={name}
            label={label}
            name={name}
            value={(form as any)[name]}
            onChange={handleChange}
            error={errors[name]}
          />
        ))}
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <ImageUploadField label="Profile Logo" name="logo" onChange={(f) => handleFileChange("logo", f)} preview={vendor.logo as string} />
        <ImageUploadField label="Cover Photo" name="coverPhoto" onChange={(f) => handleFileChange("coverPhoto", f)} preview={vendor.coverPhoto as string} />
        <ImageUploadField label="Vendor License" name="vendorLicense" onChange={(f) => handleFileChange("vendorLicense", f)} preview={vendor.vendorLicense as string} />
      </div>

      <div className="flex justify-end">
        <Button type="submit" disabled={loading}>
          {loading ? "Saving..." : "Save Changes"}
        </Button>
      </div>
    </form>
  );
}

