package profile

import (
	"fmt"
)

func GetProfile() string {
	profile := mySelf{
		name:   "Ari Ramadhan",
		photo:  "https://static.vecteezy.com/system/resources/previews/032/176/197/non_2x/business-avatar-profile-black-icon-man-of-user-symbol-in-trendy-flat-style-isolated-on-male-profile-people-diverse-face-for-social-network-or-web-vector.jpg",
		email:  "ramadhan89.ari@gmail.com",
		age:    20,
		phone:  "081268643631",
		status: "Not Married",
		education: education{
			name:  "Sekolah Tinggi Teknologi Indonesia Tanjungpinang",
			major: "Teknik Informatika"},
	}

	str := fmt.Sprintf("\nName		: %s\n Photo		:	%s\n Email		: %v\n Age		: %d\n Phone		: %s\n Status		: %s\n Education 	  %s\n 		  %s",
		profile.name, profile.photo, profile.email, profile.age, profile.phone, profile.status, profile.education.name, profile.education.major)
	return str

}
