import { z } from "zod";
import { CreateFixedAmountDiscountSchema, FixedAmountDiscountSchema } from "./fixed_discount";
import { CreatePercentageDiscountSchema, PercentageDiscountSchema } from "./percentage_discount";
import { CreatePercentageCategoryDiscountSchema, PercentageCategoryDiscountSchema } from "./percentage_category_discount";
import { DiscountCategorySchema } from "./discount_category";
import { CreateSeasonalDiscountSchema } from "./seasonal_discount";
import { CreatePointDiscountSchema } from "./point_discount";

export const PointDiscountSchema = z.object({
  id: z.number().int(),
  point: z.number(),
});

export const SeasonalDiscountSchema = z.object({
  id: z.number().int(),
  every_xthb: z.number(),
  discount_ythb: z.number(),
});

export const DiscountSchema = z.object({
  id: z.number().int(),
  fixed_amount_id: z.number().nullable(),
  fixed_amount: FixedAmountDiscountSchema.nullable(),
  percentage_id: z.number().nullable(),
  percentage: PercentageDiscountSchema.nullable(),
  percentage_category_id: z.number().nullable(),
  percentage_category: PercentageCategoryDiscountSchema.nullable(),
  seasonal_id: z.number().nullable(),
  seasonal: SeasonalDiscountSchema.nullable(),
  point_discount_id: z.number().nullable(),
  point_discount: PointDiscountSchema.nullable(),
  discount_category_id: z.number().int(),
  discount_category: DiscountCategorySchema,
});

export const CreateDiscountSchema = z.object({
  fixed_amount: CreateFixedAmountDiscountSchema.optional(),
  percentage: CreatePercentageDiscountSchema.optional(),
  percentage_category: CreatePercentageCategoryDiscountSchema.optional(),
  seasonal: CreateSeasonalDiscountSchema.optional(),
  point_discount: CreatePointDiscountSchema.optional(),
  discount_category_id: z.number().int(),
});

export type Discount = z.infer<typeof DiscountSchema>;
export type CreateDiscount = z.infer<typeof CreateDiscountSchema>;
