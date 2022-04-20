pragma solidity ^0.8.0;

contract Test{
    struct Medical {
        mapping (string => string) File;
    }

    mapping(address => Medical)  User_Medical ;
    mapping(address => string[]) public user_Medical_name;

    function AddMedical (string memory _name , string memory _url) public {
        User_Medical[msg.sender].File[_name] = _url;
        user_Medical_name[msg.sender].push(_name);
    }

    function PreviewMedicalFile (string memory _name) public view returns(string memory){
        return User_Medical[msg.sender].File[_name];
    }
}