import { z } from 'zod';
import { CartSchema } from './cart';

export const UserSchema = z.object({
  id: z.number(),
  name: z.string().optional(),
  username: z.string().optional(),
  password: z.string().optional(),
  points: z.number().default(0),
  cart: CartSchema
});

export type User = z.infer<typeof UserSchema>;