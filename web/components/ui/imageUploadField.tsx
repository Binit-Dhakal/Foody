'use client';
import Image from "next/image";
import { useState } from "react";
import { Field } from "@/components/ui/field";
import { Input } from "@/components/ui/input";

interface ImageUploadFieldProps {
  label: string;
  name: string;
  onChange: (file: File | null) => void;
  preview?: string | null;
}

export default function ImageUploadField({ label, name, onChange, preview }: ImageUploadFieldProps) {
  const [imagePreview, setImagePreview] = useState(preview || null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0] || null;
    setImagePreview(file ? URL.createObjectURL(file) : null);
    onChange(file);
  };

  return (
    <Field className="flex flex-col space-y-2">
      <label className="text-sm font-medium">{label}</label>
      <Input type="file" accept="image/*" name={name} onChange={handleFileChange} />
      {imagePreview && (
        <div className="relative w-40 h-40 mt-2 rounded-lg overflow-hidden border">
          <Image src={imagePreview} alt={label} fill className="object-cover" />
        </div>
      )}
    </Field>
  );
}

