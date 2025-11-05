'use client';
import { Field } from "@/components/ui/field";
import { Input } from "@/components/ui/input";

interface InputFieldProps {
  label: string;
  name: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  error?: string;
  type?: string;
}

export default function InputField({ label, name, value, onChange, error, type = "text" }: InputFieldProps) {
  return (
    <Field className="flex flex-col space-y-2">
      <label className="text-sm font-medium">{label}</label>
      <Input type={type} name={name} value={value} onChange={onChange} />
      {error && <p className="text-red-500 text-xs">{error}</p>}
    </Field>
  );
}

