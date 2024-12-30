import { z } from 'zod';
import { CartItemSchema } from './cart_item';

export const CartSchema = z.object({
  id: z.number(),
  user_id: z.number(),
  user: z.object({}).optional(),
  items: z.array(CartItemSchema),
});

export const CartDetailSchema = z.object({
  cart: CartSchema,
  total_price: z.number(),
  user_points_used: z.number(),
});

export type Cart = z.infer<typeof CartSchema>;
export type CartDetail = z.infer<typeof CartDetailSchema>;
