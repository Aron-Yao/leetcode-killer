int findDuplicate(int* nums, int numsSize) {
    int slow, fast, meet;
    
    slow = nums[0];
    fast = nums[nums[0]];
    
    while (slow != fast) {
        slow = nums[slow];
        fast = nums[nums[fast]];
    }
    
    meet = 0;
    while (meet != slow) {
        slow = nums[slow];
        meet = nums[meet];
    }
    
    return meet;
}
