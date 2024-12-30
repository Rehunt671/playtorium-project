
import {z} from 'zod';

export const PointDiscountSchema = z.object({
    id: z.number().int(),
    point: z.number(),
});

export const CreatePointDiscountSchema = z.object({
    point: z.number(),
});

export type PointDiscount = z.infer<typeof PointDiscountSchema>;
export type CreatePointDiscount = z.infer<typeof CreatePointDiscountSchema>;
