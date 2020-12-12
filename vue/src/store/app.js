module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "apply", fields: ["jobid", "tags", "datapath", ] },
		{ type: "joblist", fields: ["jobcreator", "tags", "limit", "members", ] },
  ],
};
