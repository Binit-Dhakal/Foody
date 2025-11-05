import { z } from 'zod';

export const vendorSchema = z.object({
  name: z.string().min(2, "Name too short"),
  address: z.string().min(3, "Address required"),
  country: z.string().optional(),
  state: z.string().optional(),
  city: z.string().optional(),
  pinCode: z.string().optional(),
  latitude: z.string().optional(),
  longitude: z.string().optional(),
  logo: z.any().optional(),
  coverPhoto: z.any().optional(),
  vendorLicense: z.any().optional(),
});

export type VendorFormData = z.infer<typeof vendorSchema>;
