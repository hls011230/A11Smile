pragma solidity ^0.8.0;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol";  

contract UploadMedicalrecords 
{

 //给合约转钱事件
event Deposit(address indexed sender, uint amount, uint balance);

//用户 
struct user_Users {
    address user;
    bool exit;
}

//医生地址
struct Doctor{
    address doctor;
    bool exist;
}

//代币地址
address ercaddress;
//上传医疗数据信息时间
uint public lastUpdateTime;

//审核医疗数据时间
uint public examineTime;

//病人上传医疗数据信息
struct user_Medicalinformation{
    address user;
    string[] PictureRoute;
    address doctor;
}

//医生审核
struct gainer_examine{
    address doctor;
    address user;
    string pRoute;
    bool  whether;
}

//modified上传时间
modifier user_upTime(address _user){
    lastUpdateTime = block.timestamp;

    _;
}

//modified审核时间
modifier gainer_examineTime(address doctor){
    examineTime = block.timestamp;

    _;
}

//给合约转钱
receive() external payable {
        emit Deposit(msg.sender, msg.value, address(this).balance);
    }

//mapping对应用户
mapping (address => user_Users)  private Ausers;

//mapping对应医生
mapping(address => Doctor) private Adoctor;

//mapping用户对应上传医疗数据信息
mapping (address => user_Medicalinformation) private userAmedicalInformation;

//mapping医生对应查询医疗数据信息
mapping(address => user_Medicalinformation) private gainAmedicalInformation;

//mapping医生审核自己的病历信息
mapping(address => gainer_examine ) private gainerExamine;
mapping(string => user_Medicalinformation ) private gainerExamineString;

//mapping查询医疗数据是否通过
mapping(string => gainer_examine ) private userExamineString;

//设置用户
function user_Adduser(address _user) public{
       Ausers[_user].user= _user;
       Ausers[_user].exit= true;
}

//设置医生
function gainer_setDoctor(address doctor) public{
    Adoctor[doctor].doctor=doctor;
    Adoctor[doctor].exist=true;
}

//用户上传医疗信息
function user_AddMedicalInformation(string memory Proute,address _doctor) public user_upTime(msg.sender) {
           require(user_IsUser(msg.sender),"patient no exist");
           require(gainer_isDoctor(_doctor),"patient no exist");
           userAmedicalInformation[msg.sender].user=msg.sender;
           userAmedicalInformation[msg.sender].PictureRoute.push(Proute);
           userAmedicalInformation[_doctor].doctor=_doctor;
}
    
//查询用户是否存在
function user_IsUser(address _user) public view returns(bool){
   return Ausers[_user].exit;
}

//查询医生是否存在
function gainer_isDoctor(address doctor)public view returns(bool){
   return Adoctor[doctor].exist;
}

//病人查看病人上传的病历
function user_ViewMedicalRecords(address _user) public view returns(string[] memory,address doctor){
    require(user_IsUser(_user),"patient no exist");
    return (userAmedicalInformation[_user].PictureRoute,userAmedicalInformation[_user].doctor);
}

//医生查看病人上传给自己的医疗信息
function  gainer_doctorsee(address doctor) public view returns(address,string[] memory,address){
    return (gainAmedicalInformation[doctor].user,gainAmedicalInformation[doctor].PictureRoute,gainAmedicalInformation[doctor].doctor);
}

//设置代币地址
function setErc(address ercaddress_) public{
         ercaddress=ercaddress_;
}

//医生审核,奖励
function gainer_whether(address doctor,string memory PictureRoute,bool whether,uint ercnum_) public gainer_examineTime(doctor){
                gainerExamine[doctor].doctor=doctor;
                gainerExamine[doctor].user=gainerExamineString[PictureRoute].user;
                gainerExamine[doctor].pRoute=PictureRoute;
                require(doctor==gainerExamineString[PictureRoute].doctor,"The medical data does not belong to the doctor");
              if(whether==true){
                  gainerExamine[doctor].whether = true;
                _transfer2other(gainerExamineString[PictureRoute].doctor,gainerExamineString[PictureRoute].user,ercnum_);
              }else{
                  revert("Failed to pass the review");
              }

}


//查看医疗数据是否通过审核
function  seeMedicaldata(string memory PictureRoute) public view returns(string memory,bool){
        return(PictureRoute,userExamineString[PictureRoute].whether);
}

//医生奖励
    function _transfer2other(address from_,address to_ ,uint ercnum_) public {
        IERC20(ercaddress).transferFrom(from_,to_,ercnum_);
   
   }
}