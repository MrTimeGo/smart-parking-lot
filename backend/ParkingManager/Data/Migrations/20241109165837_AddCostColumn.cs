﻿using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace ParkingManager.Data.Migrations
{
    /// <inheritdoc />
    public partial class AddCostColumn : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<decimal>(
                name: "Cost",
                table: "ActionLogs",
                type: "numeric",
                nullable: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "Cost",
                table: "ActionLogs");
        }
    }
}
