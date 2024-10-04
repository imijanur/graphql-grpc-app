// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("UserAddresses", testUserAddresses)
	t.Run("UserContacts", testUserContacts)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesDelete)
	t.Run("UserContacts", testUserContactsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesQueryDeleteAll)
	t.Run("UserContacts", testUserContactsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesSliceDeleteAll)
	t.Run("UserContacts", testUserContactsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesExists)
	t.Run("UserContacts", testUserContactsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesFind)
	t.Run("UserContacts", testUserContactsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesBind)
	t.Run("UserContacts", testUserContactsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesOne)
	t.Run("UserContacts", testUserContactsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesAll)
	t.Run("UserContacts", testUserContactsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesCount)
	t.Run("UserContacts", testUserContactsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesHooks)
	t.Run("UserContacts", testUserContactsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesInsert)
	t.Run("UserAddresses", testUserAddressesInsertWhitelist)
	t.Run("UserContacts", testUserContactsInsert)
	t.Run("UserContacts", testUserContactsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesReload)
	t.Run("UserContacts", testUserContactsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesReloadAll)
	t.Run("UserContacts", testUserContactsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesSelect)
	t.Run("UserContacts", testUserContactsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesUpdate)
	t.Run("UserContacts", testUserContactsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("UserAddresses", testUserAddressesSliceUpdateAll)
	t.Run("UserContacts", testUserContactsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
