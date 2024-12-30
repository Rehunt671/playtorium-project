import { useState } from "react";
import { CreateDiscount } from "@/types/discount";
import { useMutationCreateDiscount } from "@/query/discount";

interface PointDiscountFormProps {
  selectedDiscountCategoryId: number; 
}

const PointDiscountForm = ({ selectedDiscountCategoryId }: PointDiscountFormProps) => {
  const [point, setPoint] = useState<number>(0);
  const { mutateAsync, isPending, error } = useMutationCreateDiscount();

  const onSubmitDiscount = async (e: React.FormEvent) => {
    e.preventDefault();

    const discountData: CreateDiscount = {
      point_discount: {
        point: point,
      },
      discount_category_id: selectedDiscountCategoryId,
    };

    try {
      await mutateAsync(discountData);
    } catch (err) {
      console.error("Error creating discount", err);
    }
  };

  return (
    <form onSubmit={onSubmitDiscount}>
      <h2 className="text-xl font-bold text-gray-800 mb-4">Point Discount</h2>

      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Points
        </label>
        <input
          type="number"
          name="points"
          value={point}
          onChange={(e) => setPoint(Number(e.target.value))}
          placeholder="Enter points"
          className="block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      {error && <p className="text-red-500 text-sm mb-4">{error.message}</p>} {/* Display error message */}

      <button
        type="submit"
        className="w-full bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        disabled={isPending}
      >
        {isPending ? "Creating..." : "Apply Points"}
      </button>
    </form>
  );
};

export default PointDiscountForm;
