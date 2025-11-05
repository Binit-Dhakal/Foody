export default function DashboardPage() {
    return (
        <div className="space-y-6">
            <h2 className="text-2xl font-bold text-gray-800">Your Dashboard</h2>

            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <div className="p-4 bg-white rounded-lg shadow-sm border border-gray-100">
                    <h3 className="font-semibold text-gray-700">Last Order</h3>
                    <p className="text-gray-500 text-sm mt-1">Pizza from Domino’s • Delivered</p>
                </div>

                <div className="p-4 bg-white rounded-lg shadow-sm border border-gray-100">
                    <h3 className="font-semibold text-gray-700">Favorite Restaurant</h3>
                    <p className="text-gray-500 text-sm mt-1">Pizza Hut</p>
                </div>

                <div className="p-4 bg-white rounded-lg shadow-sm border border-gray-100">
                    <h3 className="font-semibold text-gray-700">Total Orders</h3>
                    <p className="text-gray-500 text-sm mt-1">23 Orders completed</p>
                </div>
            </div>
        </div>
    );
}


