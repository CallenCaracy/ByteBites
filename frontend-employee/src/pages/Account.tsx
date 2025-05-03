import { useParams } from "react-router-dom";
import { useMutation, useQuery } from "@apollo/client";
import React, { useState } from "react";
import { GET_USER_BY_ID, UPDATE_USER  } from "../graphql/Userqueries";
import Navbar from "../components/NavBar";
import { supabase } from '../utils/supabaseClient';

const Account: React.FC = () => {
    const { userId } = useParams();
    const { data, loading, error } = useQuery(GET_USER_BY_ID, {
        variables: { id: userId },
    });
    const [updateUser] = useMutation(UPDATE_USER);
    const [ isEditing, setIsEditing] = useState(false);
    const [formData, setFormData] = useState({
        firstName: "",
        lastName: "",
        email: "",
        address: "",
        phone: "",
        isActive: "active",
        gender: "",
        userType: "",
        pfp: "",
        birthDate: "",
    });
    const [formErrors, setFormErrors] = useState({
        phone: "",
    });

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error fetching user data: {error.message}</p>;
    if (!data || !data.getUserById) return <p>No user data found</p>;

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: value,
        }));

        let errors = { ...formErrors };

        if (name === "phone") {
            if (!/^\d*$/.test(value)) {
                errors.phone = "Phone number must contain only digits.";
            } else if (value.length > 11) {
                errors.phone = "Phone number cannot exceed 11 digits.";
            } else if (value.length < 11) {
                errors.phone = "Phone number cannot be less than 11 digits.";
            }else {
                errors.phone = "";
            }
        }

        setFormErrors(errors);
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        const formattedBirthDate = formData.birthDate
        ? new Date(formData.birthDate).toISOString().split("T")[0]
        : "";

        try {
            await updateUser({
                variables: {
                    id: userId,
                    input: {
                        firstName: formData.firstName,
                        lastName: formData.lastName,
                        address: formData.address,
                        phone: formData.phone,
                        birthDate: formattedBirthDate,
                        userType: formData.userType,
                        gender: formData.gender,
                        isActive: formData.isActive,
                    },
                },
            });
            setIsEditing(false);
        } catch (error) {
            console.error("Error updating user:", error);
        }
    };

    const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (!file) return;
    
        const oldPfpUrl = formData.pfp;
        let oldPfpPath: string | null = null;
    
        if (oldPfpUrl) {
            try {
                const url = new URL(oldPfpUrl);
                const pathname = url.pathname;
                const match = pathname.match(/\/storage\/v1\/object\/public\/pictures\/(.+)/);

                if (match) {
                    oldPfpPath = match[1];
                }
            } catch (err) {
                console.warn('Failed to parse oldPfpUrl:', err);
            }
        }
    
        const fileExt = file.name.split('.').pop();
        const fileName = `${userId}-${Date.now()}.${fileExt}`;
        const filePath = `pfp/${fileName}`;
    
        if (oldPfpPath && oldPfpPath !== 'pfp/defualtpic.jpg') {
            const { error: deleteError } = await supabase
                .storage
                .from('pictures')
                .remove([oldPfpPath]);
                const { data: items, error: listError } = await supabase
                .storage
                .from('pictures')
                .list('pfp', { limit: 100 });
                console.log("🗂 Remaining in pfp/:", items, listError);
    
            if
            (deleteError) {
                console.error('Error deleting old file:', deleteError.message);
            } else {
                console.log('Old file deleted successfully');
            }
        }
    
        const { error: uploadError } = await supabase
            .storage
            .from('pictures')
            .upload(filePath, file);
    
        if (uploadError) {
            console.error('Error uploading file:', uploadError.message);
            return;
        }
    
        const { data: publicUrlData } = supabase
            .storage
            .from('pictures')
            .getPublicUrl(filePath);
    
        const publicUrl = publicUrlData.publicUrl;
    
        setFormData((prev) => ({
            ...prev,
            pfp: publicUrl,
        }));
    
        try {
            await updateUser({
                variables: {
                    id: userId,
                    input: {
                        pfp: publicUrl,
                    },
                },
            });
        } catch (error) {
            console.error("Error updating profile picture:", error);
        }
    };    

    const handleEditClick = () => {
        setIsEditing(true);

        const formattedBirthDate = data.getUserById.birthDate
        ? new Date(data.getUserById.birthDate).toISOString().split("T")[0]
        : "";

        setFormData({
            ...formData,
            firstName: data.getUserById.firstName,
            lastName: data.getUserById.lastName,
            email: data.getUserById.email,
            address: data.getUserById.address || "",
            phone: data.getUserById.phone || "",
            isActive: data.getUserById.isActive === "active" ? "active" : "inactive",
            gender: data.getUserById.gender || "",
            userType: data.getUserById.userType || "",
            pfp: data.getUserById.pfp || "",
            birthDate: formattedBirthDate || "",
        });
    };

    return (
        <div>
            <Navbar />
            {/* <h1 className="text-3xl font-semibold text-gray-800 mb-6">Account Profile</h1> */}
            <div className="bg-white p-6 rounded-lg shadow-md flex flex-col md:flex-row gap-8">
                <div className="flex flex-col items-center md:w-1/4">
                <h2 className="text-xl font-semibold text-gray-700">User Info</h2>
                    <div className="flex justify-center items-center mb-4">
                        <img
                            src={formData.pfp || data?.getUserById?.pfp}
                            alt="Profile Picture"
                            className="w-50 h-50 rounded-full object-cover"
                        />
                    </div>
                    {isEditing && (
                        <div className="md:w-2/3 w-full">
                            <label className="block text-sm font-medium text-gray-700 mb-1">Change Profile Picture</label>
                            <input
                                title="Pfp"
                                type="file"
                                accept="image/*"
                                onChange={handleFileChange}
                                className="text-black block w-full px-3 py-2 border border-gray-300 rounded-md"
                            />
                        </div>
                    )}
                </div>
                <div>
                    {isEditing ? (
                        <form onSubmit={handleSubmit}>
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                                <div>
                                    <label className="block text-sm font-medium text-gray-700">First Name</label>
                                    <input
                                        title="First Name"
                                        type="text"
                                        name="firstName"
                                        value={formData.firstName}
                                        onChange={handleInputChange}  // Use the same handler here
                                        className="text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
                                    />
                                </div>
                                <div>
                                    <label className="block text-sm font-medium text-gray-700">Last Name</label>
                                    <input
                                        title="Last Name"
                                        type="text"
                                        name="lastName"
                                        value={formData.lastName}
                                        onChange={handleInputChange}
                                        className="text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
                                    />
                                </div>
                                <div>
                                    <label className="block text-sm font-medium text-gray-700">Address</label>
                                    <input
                                        title="Address"
                                        type="text"
                                        name="address"
                                        value={formData.address}
                                        onChange={handleInputChange}
                                        className="text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
                                    />
                                </div>
                                <div>
                                    <label className="block text-sm font-medium text-gray-700">Phone</label>
                                    <input
                                        title="Phone"
                                        type="text"
                                        name="phone"
                                        value={formData.phone}
                                        onChange={handleInputChange}
                                        className="text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
                                    />
                                    {formErrors.phone && <p className="text-red-500 text-sm">{formErrors.phone}</p>}
                                </div>
                                <div>
                                    <label className="block text-sm font-medium text-gray-700">Birth Date</label>
                                    <input
                                        title="Birth Date"
                                        type="date"
                                        name="birthDate"
                                        value={formData.birthDate}
                                        onChange={handleInputChange}
                                        className="text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
                                    />
                                </div>
                                <div>
                                    <label className="block text-sm font-medium text-gray-700">Gender</label>
                                    <select
                                        title="Gender"
                                        name="gender"
                                        value={formData.gender}
                                        onChange={handleInputChange}  // Same handler here for select
                                        className="text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
                                    >
                                        <option value="male">Male</option>
                                        <option value="female">Female</option>
                                        <option value="other">Other</option>
                                    </select>
                                </div>
                                <div className="mb-4">
                                <button
                                    type="submit"
                                    className={`px-4 py-2 rounded-md ${
                                        Object.values(formErrors).some((error) => error) 
                                            ? 'bg-gray-400 cursor-not-allowed' 
                                            : 'bg-blue-500 hover:bg-blue-600'
                                    }`}
                                    disabled={Object.values(formErrors).some((error) => error)}
                                >
                                    Update
                                </button>
                                    <button
                                        type="button"
                                        onClick={() => setIsEditing(false)}
                                        className="px-4 py-2 bg-gray-500 rounded-md ml-2">
                                        Close
                                    </button>
                                </div>
                            </div>
                        </form>
                    ) : (
                        <>
                            <p className="text-black"><strong>Account Is:</strong> {data?.getUserById?.isActive || "No status info available"}</p>
                            <p className="text-black"><strong>Name:</strong> {data?.getUserById?.firstName || "No name available"} {data?.getUserById?.lastName || "No last name available"}</p>
                            <p className="text-black"><strong>Email:</strong> {data?.getUserById?.email || "No email available"}</p>
                            <p className="text-black">
                            <strong>Birth Date:</strong> {data?.getUserById?.birthDate 
                                ? new Date(data.getUserById.birthDate).toISOString().split("T")[0]
                                : "No birthDate available"}
                            </p>
                            <p className="text-black"><strong>Gender:</strong> {data?.getUserById?.gender || "No gender available"}</p>
                            <p className="text-black"><strong>User Type:</strong> {data?.getUserById?.userType || "No user type available"}</p>
                            <p className="text-black"><strong>Role:</strong> {data?.getUserById?.role || "No role available"}</p>
                            <p className="text-black"><strong>Address:</strong> {data?.getUserById?.address || "No address available"}</p>
                            <p className="text-black"><strong>Phone:</strong> {data?.getUserById?.phone || "No phone number available"}</p>
                            <button onClick={handleEditClick} className="mt-2 text-white">Update Profile</button>
                        </>
                    )}
                </div>
            </div>
        </div>
    );
};

export default Account;
