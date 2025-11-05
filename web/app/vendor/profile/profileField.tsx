"use client";
import { useState } from "react";

export default function ProfileField({ label, value, field }: { label: string; value: string; field: string }) {
  const [edit, setEdit] = useState(false);
  const [newValue, setNewValue] = useState(value);

  const handleSave = async () => {
    await fetch("/api/vendor/profile", {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ [field]: newValue }),
    });
    setEdit(false);
  };

  return (
    <div className="flex justify-between items-center border-b py-2">
      <span className="font-medium">{label}</span>
      {edit ? (
        <div className="flex space-x-2">
          <input className="border p-1" value={newValue} onChange={(e) => setNewValue(e.target.value)} />
          <button onClick={handleSave} className="text-green-600">Save</button>
          <button onClick={() => setEdit(false)} className="text-gray-500">Cancel</button>
        </div>
      ) : (
        <div className="flex items-center space-x-2">
          <span>{value || "—"}</span>
          <button onClick={() => setEdit(true)} className="text-blue-500 text-sm">Edit</button>
        </div>
      )}
    </div>
  );
}

