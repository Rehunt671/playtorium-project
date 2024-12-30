import { Discount } from '@/types/discount';
import React from 'react';

type DiscountProps = {
  discount: Discount;
  onRemove?: () => Promise<void>;
};

const DiscountCard: React.FC<DiscountProps> = ({ discount, onRemove }) => {
  return (
    <div className="mt-6 p-6 bg-white rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-300 ease-in-out">
      <div className="mb-4">
        {discount.percentage && (
          <div className="mb-2">
            <h3 className="text-xl font-semibold text-gray-800">Percentage Discount</h3>
            <p className="text-gray-600">Discount: {discount.percentage.percentage}%</p>
          </div>
        )}
        {discount.fixed_amount && (
          <div className="mb-2">
            <h3 className="text-xl font-semibold text-gray-800">Fixed Amount Discount</h3>
            <p className="text-gray-600">Discount: ${discount.fixed_amount.amount.toFixed(2)}</p>
          </div>
        )}
        {discount.seasonal && (
          <div className="mb-2">
            <h3 className="text-xl font-semibold text-gray-800">Seasonal Discount</h3>
            <p className="text-gray-600">
              Every {discount.seasonal.every_xthb} Discount: {discount.seasonal.discount_ythb}
            </p>
          </div>
        )}
        {discount.point_discount && (
          <div className="mb-2">
            <h3 className="text-xl font-semibold text-gray-800">Points Discount</h3>
            <p className="text-gray-600">Discount: {discount.point_discount.point} points</p>
          </div>
        )}
        {discount.percentage_category && (
          <div className="mb-2">
            <h3 className="text-xl font-semibold text-gray-800">Percentage Category Discount</h3>
            <p className="text-gray-600">Category: {discount.percentage_category.item_category.name}</p>
            <p className="text-gray-600">Discount: {discount.percentage_category.percentage}%</p>
          </div>
        )}
      </div>

      {/* Remove Button */}
      {onRemove && (
        <div className="mt-4 flex justify-end">
          <button
            onClick={onRemove}
            className="px-5 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg shadow-md hover:shadow-lg transition-all duration-200 ease-in-out"
          >
            Remove
          </button>
        </div>
      )}
    </div>
  );
};

export default DiscountCard;
